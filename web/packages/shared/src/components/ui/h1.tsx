import clsx from 'clsx';

interface H1Props extends React.HTMLAttributes<HTMLHeadingElement> {
	children: React.ReactNode;
	className?: string;
}

export function H1({ children, className, ...props }: H1Props) {
	return (
		<h1
			className={clsx(
				'scroll-m-20 text-center text-4xl font-extrabold tracking-tight text-balance',
				className,
			)}
			{...props}
		>
			{children}
		</h1>
	);
}
