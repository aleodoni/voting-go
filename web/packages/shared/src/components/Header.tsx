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
		<header className="w-full rounded-xl border bg-background shadow-sm px-6 py-4">
			<div className="flex items-center justify-between">
				<div className="flex items-center gap-4">
					<Link
						to="/"
						title="Página inicial"
						className="flex items-center justify-center"
					>
						<img
							alt="Logo Câmara Municipal de Curitiba"
							src={brasaoPng}
							className="h-auto w-14 lg:w-16"
						/>
					</Link>

					<div className="flex flex-col">
						<p className="text-xl lg:text-3xl font-bold leading-tight">
							Câmara Municipal de Curitiba
						</p>

						<p className="text-base lg:text-xl font-semibold">
							Sistema de Votação
						</p>

						<p className="text-sm opacity-70">{subtitulo}</p>
					</div>
				</div>

				<nav className="flex items-center gap-3">
					<ThemeTogle setTheme={setTheme} />
					<ButtonLogout logout={logout} />
				</nav>
			</div>
		</header>
	);
}
