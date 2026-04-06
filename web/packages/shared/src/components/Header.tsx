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
		<header className="w-full rounded-xl border bg-background shadow-sm px-4 py-3 sm:px-6 sm:py-4">
			<div className="flex items-center justify-between gap-2">
				<div className="flex items-center gap-2 sm:gap-4 min-w-0">
					<Link
						to="/"
						title="Página inicial"
						className="flex items-center justify-center shrink-0"
					>
						<img
							alt="Logo Câmara Municipal de Curitiba"
							src={brasaoPng}
							className="h-auto w-10 sm:w-14 lg:w-16"
						/>
					</Link>

					<div className="flex flex-col min-w-0">
						<p className="text-sm sm:text-xl lg:text-3xl font-bold leading-tight truncate">
							Câmara Municipal de Curitiba
						</p>

						<p className="text-xs sm:text-base lg:text-xl font-semibold truncate">
							Sistema de Votação
						</p>

						<p className="text-xs sm:text-sm opacity-70 truncate">
							{subtitulo}
						</p>
					</div>
				</div>

				<nav className="flex items-center gap-2 sm:gap-3 shrink-0">
					<ThemeTogle setTheme={setTheme} />
					<ButtonLogout logout={logout} />
				</nav>
			</div>
		</header>
	);
}
