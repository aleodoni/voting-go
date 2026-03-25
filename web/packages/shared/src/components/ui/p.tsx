import clsx from 'clsx';

interface PProps extends React.HTMLAttributes<HTMLParagraphElement> {
	children: React.ReactNode;
	className?: string;
}

export function P({ children, className, ...props }: PProps) {
	return (
		<p
			className={clsx('leading-7 [&:not(:first-child)]:mt-6', className)}
			{...props}
		>
			{children}
		</p>
	);
}
