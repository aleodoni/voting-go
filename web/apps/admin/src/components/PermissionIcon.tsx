import { ThumbsDown, ThumbsUp } from 'lucide-react';

type PermissionIconProps = {
	status: boolean;
};

export function PermissionIcon({ status }: PermissionIconProps) {
	if (status) {
		return (
			<ThumbsUp className="flex items-center justify-center w-5 h-5 text-green-500" />
		);
	} else {
		return (
			<ThumbsDown className="flex items-center justify-center w-5 h-5 text-red-500" />
		);
	}
}
