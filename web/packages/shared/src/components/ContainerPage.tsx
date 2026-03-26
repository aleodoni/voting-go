import clsx from 'clsx';

interface ContainerPageProps extends React.HTMLAttributes<HTMLDivElement> {
	children: React.ReactNode;
	className?: string;
}

export function ContainerPage({
	children,
	className,
	...props
}: ContainerPageProps) {
	return (
		<div
			className={clsx(
				'flex flex-col bg-background border rounded-xl p-8 my-8 items-start h-full w-full',
				className,
			)}
			{...props}
		>
			{children}
		</div>
	);
}
