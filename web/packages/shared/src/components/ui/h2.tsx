import clsx from 'clsx';

interface H2Props extends React.HTMLAttributes<HTMLHeadingElement> {
	children: React.ReactNode;
	className?: string;
}

export function H2({ children, className, ...props }: H2Props) {
	return (
		<h2
			className={clsx(
				'scroll-m-20 border-b pb-2 text-3xl font-semibold tracking-tight first:mt-0 w-full',
				className,
			)}
			{...props}
		>
			{children}
		</h2>
	);
}
