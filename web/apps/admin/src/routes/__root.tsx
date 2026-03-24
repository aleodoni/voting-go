import { createRootRoute, Outlet } from '@tanstack/react-router';

export const Route = createRootRoute({
	component: RootComponent,
});

function RootComponent() {
	return (
		<div className="flex w-full h-full">
			<Outlet />
		</div>
	);
}
