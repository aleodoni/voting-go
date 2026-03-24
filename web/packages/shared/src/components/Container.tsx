import clsx from 'clsx';

interface ContainerProps extends React.HTMLAttributes<HTMLDivElement> {
	children: React.ReactNode;
	className?: string;
}

export function Container({ children, className, ...props }: ContainerProps) {
	return (
		<div
			className={clsx(
				'flex flex-col w-full h-screen max-w-6xl mx-auto',
				className,
			)}
			{...props}
		>
			{children}
		</div>
	);
}
