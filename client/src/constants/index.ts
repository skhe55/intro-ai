import type { TSquare } from '$lib/types';

//api
export const DEFAULT_API_PATH = 'http://localhost:3000';
export const JSON_BODY_HEADERS = {
	headers: {
		'Content-Type': 'application/json'
	}
};

export const initialSquareCoordinates: Omit<TSquare, 'id'> = {
	x_top: [0, 0],
	x_bottom: [0, 0],
	y_top: [0, 0],
	y_bottom: [0, 0]
};
