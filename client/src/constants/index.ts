import type { TSquare } from '$lib/types';

export const initialSquareCoordinates: Omit<TSquare, 'id'> = {
	x_top: [0, 0],
	x_bottom: [0, 0],
	y_top: [0, 0],
	y_bottom: [0, 0]
};
