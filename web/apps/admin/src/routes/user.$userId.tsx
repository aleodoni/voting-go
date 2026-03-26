import { zodResolver } from '@hookform/resolvers/zod';
import { createFileRoute } from '@tanstack/react-router';
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
	Header,
	Input,
	Switch,
	useAuth,
	useTheme,
} from '@voting/shared';
import { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { z } from 'zod';
import { useUpdateUser } from '@/hooks/useUpdateUser';
import { useUser } from '@/hooks/useUser';

export const Route = createFileRoute('/user/$userId')({
	component: UserDetail,
});

const schema = z.object({
	nome: z.string().min(1, 'Nome obrigatório'),
	email: z.string().email('Email inválido'),
	nome_fantasia: z.string().min(1, 'Nome fantasia obrigatório'),
	ativo: z.boolean(),
	pode_administrar: z.boolean(),
	pode_votar: z.boolean(),
});

type FormData = z.infer<typeof schema>;

function UserDetail() {
	const { userId } = Route.useParams();

	const { data: user, isLoading } = useUser(userId);
	const { updateUser, isPending } = useUpdateUser();

	const { setTheme } = useTheme();
	const { logout } = useAuth();

	const form = useForm<FormData>({
		resolver: zodResolver(schema),
		defaultValues: {
			nome: '',
			email: '',
			nome_fantasia: '',
			ativo: false,
			pode_administrar: false,
			pode_votar: false,
		},
	});

	useEffect(() => {
		if (user) {
			form.reset({
				nome: user.nome,
				email: user.email,
				nome_fantasia: user.nome_fantasia ?? '',
				ativo: user.credencial.ativo,
				pode_administrar: user.credencial.pode_administrar,
				pode_votar: user.credencial.pode_votar,
			});
		}
	}, [user, form]);

	async function onSubmit(values: FormData) {
		await updateUser(userId, values);
	}

	if (!user) {
		return <div>Usuário não encontrado</div>;
	}

	return (
		<div className="w-full min-h-screen px-6 py-4">
			<div className="max-w-7xl mx-auto flex flex-col gap-4">
				<Header
					subtitulo="Módulo administrativo"
					logout={logout}
					setTheme={setTheme}
				/>

				<ContainerPage>
					<H2>Manutenção de usuário</H2>

					<Form {...form}>
						<form
							onSubmit={form.handleSubmit(onSubmit)}
							className="w-full space-y-4 py-6"
						>
							<FormField
								control={form.control}
								name="nome"
								render={({ field }) => (
									<FormItem>
										<FormLabel>Nome</FormLabel>
										<FormControl>
											<Input {...field} disabled={isPending} />
										</FormControl>
										<FormMessage />
									</FormItem>
								)}
							/>

							<FormField
								control={form.control}
								name="email"
								render={({ field }) => (
									<FormItem>
										<FormLabel>Email</FormLabel>
										<FormControl>
											<Input {...field} disabled={isPending} />
										</FormControl>
										<FormMessage />
									</FormItem>
								)}
							/>

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
								<Button type="submit" disabled={isPending}>
									{isPending ? (
										<span className="animate-spin">⏳</span>
									) : (
										'Salvar'
									)}
								</Button>
							</div>
						</form>
					</Form>
				</ContainerPage>
			</div>
		</div>
	);
}
