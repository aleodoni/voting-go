import { Link } from '@tanstack/react-router';
import brasaoPng from '../assets/brasao.png';
import { ButtonLogout } from './ButtonLogout';
import { ThemeTogle } from './ThemeTogle';

interface HeaderProps {
	subtitulo: string;
	logout: () => void;
	setTheme: (theme: 'light' | 'dark' | 'system') => void;
}
export function Header({ logout, subtitulo, setTheme }: HeaderProps) {
	return (
		<div className="flex items-center w-full justify-between border-b py-2">
			<div className="flex h-full lg:w-2/3 py-3">
				<Link
					className="ml-4 flex items-center justify-center text-2xl"
					to="/"
					title="Página inicial"
				>
					<img
						alt="Logo Câmara Municipal de Curitiba"
						src={brasaoPng}
						className="flex h-auto w-[50px] lg:w-[80px] items-center"
					/>
				</Link>
				<div className="ml-4 flex flex-col justify-center gap-2">
					<p className="flex text-lg lg:text-4xl font-bold">
						Câmara Municipal de Curitiba
					</p>
					<div className="flex flex-col">
						<p className="text-base lg:text-2xl font-bold">
							Sistema de Votação
						</p>
						<p className="text-sm lg:text-md opacity-80">{subtitulo}</p>
					</div>
				</div>
			</div>
			<nav className=" flex lg:w-1/3 items-center justify-end gap-8">
				<ButtonLogout logout={logout} />
				<ThemeTogle setTheme={setTheme} />
			</nav>
		</div>
	);
}
