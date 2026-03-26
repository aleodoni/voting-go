import { zodResolver } from '@hookform/resolvers/zod';
import {
	Button,
	Form,
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
	getApi,
	Input,
	User,
	useAuth,
} from '@voting/shared';
import { useState } from 'react';
import { useForm } from 'react-hook-form';
import { toast } from 'react-hot-toast';
import { z } from 'zod';

type FormUserInfoProps = {
	userInfo: User;
};

export function FormUserInfo({ userInfo }: FormUserInfoProps) {
	const [isLoading, setIsLoading] = useState(false);

	const { refreshUser } = useAuth();

	const form = useForm<z.infer<typeof formUserinfoSchema>>({
		shouldUnregister: true,
		resolver: zodResolver(formUserinfoSchema),
		defaultValues: {
			display_name: userInfo.nome_fantasia || '',
		},
	});

	const onSubmit = async (values: z.infer<typeof formUserinfoSchema>) => {
		setIsLoading(true);

		try {
			const { data } = await getApi().put('/usuarios/fantasia', {
				user_id: userInfo.id,
				display_name: values.display_name,
			});

			toast.success(data.message || 'Nome atualizado com sucesso!');
			await refreshUser();
		} catch (error) {
			console.error(error);
			toast.error('Erro ao atualizar o nome');
		} finally {
			setIsLoading(false);
		}
	};

	return (
		<Form {...form}>
			<form onSubmit={form.handleSubmit(onSubmit)} className="w-full space-y-4">
				<FormField
					control={form.control}
					name="display_name"
					render={({ field }) => (
						<FormItem>
							<FormLabel>Nome no sistema de votação</FormLabel>
							<FormControl>
								<Input {...field} disabled={isLoading} />
							</FormControl>
							<FormMessage />
						</FormItem>
					)}
				/>

				<div className="flex w-full justify-end">
					<Button type="submit" disabled={isLoading}>
						{isLoading ? <span className="animate-spin">⏳</span> : 'Salvar'}
					</Button>
				</div>
			</form>
		</Form>
	);
}

const formUserinfoSchema = z.object({
	display_name: z.string().nonempty({ message: 'Nome no sistema obrigatório' }),
});
