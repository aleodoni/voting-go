import { zodResolver } from '@hookform/resolvers/zod';
import { createFileRoute, useNavigate } from '@tanstack/react-router';
import {
	Button,
	ContainerPage,
	Form,
	FormControl,
	FormDescription,
	FormField,
	FormItem,
	FormLabel,
	FormMessage,
	H2,
	Input,
	Switch,
} from '@voting/shared';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import toast from 'react-hot-toast';
import { z } from 'zod';
import { useUpdateUser } from '@/hooks/useUpdateUser';
import { useUser } from '@/hooks/useUser';

export const Route = createFileRoute('/user/$userId')({
	component: UserDetail,
	validateSearch: (search) =>
		z
			.object({
				returnSearch: z
					.object({
						nome: z.string().optional(),
						email: z.string().optional(),
						page: z.number().optional(),
					})
					.optional(),
			})
			.parse(search),
});

const schema = z.object({
	nome_fantasia: z.string().min(1, 'Nome fantasia obrigatório'),
	ativo: z.boolean(),
	pode_administrar: z.boolean(),
	pode_votar: z.boolean(),
});

type FormData = z.infer<typeof schema>;

function UserDetail() {
	const { userId } = Route.useParams();

	const { returnSearch } = Route.useSearch();

	const navigate = useNavigate();

	const { data: user } = useUser(userId);
	const { updateUser, isPending } = useUpdateUser();

	const form = useForm<FormData>({
		resolver: zodResolver(schema),
		defaultValues: {
			nome_fantasia: '',
			ativo: false,
			pode_administrar: false,
			pode_votar: false,
		},
	});

	useEffect(() => {
		if (user) {
			form.reset({
				nome_fantasia: user.nome_fantasia ?? '',
				ativo: user.credencial.ativo,
				pode_administrar: user.credencial.pode_administrar,
				pode_votar: user.credencial.pode_votar,
			});
		}
	}, [user, form]);

	async function onSubmit(values: FormData) {
		await updateUser(userId, values);
		toast.success('Usuário atualizado com sucesso!');
		navigate({ to: '/manage-users', search: returnSearch ?? {} });
	}

	if (!user) {
		return <div>Usuário não encontrado</div>;
	}

	return (
		<ContainerPage>
			<H2>Manutenção de usuário</H2>

			<Form {...form}>
				<form
					onSubmit={form.handleSubmit(onSubmit)}
					className="w-full space-y-4 py-6"
				>
					<div className="space-y-2">
						<p className="text-sm font-medium">Nome</p>
						<p className="text-sm text-muted-foreground">{user.nome}</p>
					</div>

					<div className="space-y-2">
						<p className="text-sm font-medium">Email</p>
						<p className="text-sm text-muted-foreground">{user.email}</p>
					</div>

					<FormField
						control={form.control}
						name="nome_fantasia"
						render={({ field }) => (
							<FormItem>
								<FormLabel>Nome fantasia</FormLabel>
								<FormControl>
									<Input {...field} disabled={isPending} />
								</FormControl>
								<FormMessage />
							</FormItem>
						)}
					/>

					<FormField
						control={form.control}
						name="ativo"
						render={({ field }) => (
							<FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
								<div className="space-y-0.5">
									<FormLabel>Usuário ativo</FormLabel>
									<FormDescription>
										Usuário está ativo no sistema.
									</FormDescription>
								</div>
								<FormControl>
									<Switch
										checked={field.value}
										onCheckedChange={field.onChange}
									/>
								</FormControl>
								<FormMessage />
							</FormItem>
						)}
					/>

					<FormField
						control={form.control}
						name="pode_administrar"
						render={({ field }) => (
							<FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
								<div className="space-y-0.5">
									<FormLabel>Pode administrar</FormLabel>
									<FormDescription>
										Usuário pode executar ações administrativas.
									</FormDescription>
								</div>
								<FormControl>
									<Switch
										checked={field.value}
										onCheckedChange={field.onChange}
									/>
								</FormControl>
								<FormMessage />
							</FormItem>
						)}
					/>

					<FormField
						control={form.control}
						name="pode_votar"
						render={({ field }) => (
							<FormItem className="flex flex-row items-center justify-between rounded-lg border p-3 shadow-sm">
								<div className="space-y-0.5">
									<FormLabel>Pode votar</FormLabel>
									<FormDescription>
										Usuário pode participar das votações.
									</FormDescription>
								</div>
								<FormControl>
									<Switch
										checked={field.value}
										onCheckedChange={field.onChange}
									/>
								</FormControl>
								<FormMessage />
							</FormItem>
						)}
					/>

					<div className="flex w-full justify-end">
						<Button type="submit" disabled={isPending} variant={'outline'}>
							{isPending ? <span className="animate-spin">⏳</span> : 'Salvar'}
						</Button>
					</div>
				</form>
			</Form>
		</ContainerPage>
	);
}
