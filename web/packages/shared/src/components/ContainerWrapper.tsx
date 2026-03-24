import clsx from 'clsx';

interface ContainerWrapperProps extends React.HTMLAttributes<HTMLDivElement> {
	children: React.ReactNode;
	className?: string;
}

export function ContainerWrapper({
	children,
	className,
	...props
}: ContainerWrapperProps) {
	return (
		<div
			className={clsx('flex flex-col w-full min-h-screen', className)}
			{...props}
		>
			{children}
		</div>
	);
}
