import { useMatch, useNavigate } from '@tanstack/react-router';
import { Button, Checkbox, Input } from '@voting/shared';
import { useRef, useState } from 'react';
import { Route as ManageUsersRoute } from '@/routes/manage-users';

export function SearchUsers() {
	const nomeRef = useRef<HTMLInputElement>(null);
	const emailRef = useRef<HTMLInputElement>(null);

	const navigate = useNavigate();

	const match = useMatch({ from: ManageUsersRoute.id });
	const {
		nome = '',
		email = '',
		listarInativos: listarInativosDefault = false,
	} = match.search ?? {};

	const [listarInativos, setListarInativos] = useState<boolean>(
		listarInativosDefault,
	);

	function handleSearch() {
		navigate({
			to: '.',
			search: {
				nome: nomeRef.current?.value || '',
				email: emailRef.current?.value || '',
				listarInativos,
				page: 1,
			},
			replace: true,
		});
	}

	return (
		<section className="flex flex-col w-full gap-4 sm:flex-row sm:items-end">
			<div className="flex flex-col w-full gap-2">
				<p className="text-sm font-medium">Usuário</p>
				<Input
					type="text"
					ref={nomeRef}
					defaultValue={nome}
					onKeyDown={(e) => e.key === 'Enter' && handleSearch()}
				/>
			</div>

			<div className="flex flex-col w-full gap-2">
				<p className="text-sm font-medium">E-mail</p>
				<Input
					type="text"
					ref={emailRef}
					defaultValue={email}
					onKeyDown={(e) => e.key === 'Enter' && handleSearch()}
				/>
			</div>

			<div className="flex flex-col gap-1">
				<p className="text-sm font-medium invisible">Filtro</p>

				<div className="flex items-center gap-2 h-[40px]">
					<Checkbox
						checked={listarInativos}
						onCheckedChange={(value: boolean | 'indeterminate') =>
							setListarInativos(!!value)
						}
					/>
					<p className="text-sm font-medium">Inativos</p>
				</div>
			</div>

			<Button
				onClick={handleSearch}
				variant="outline"
				className="w-full sm:w-auto shrink-0"
			>
				Pesquisar
			</Button>
		</section>
	);
}
