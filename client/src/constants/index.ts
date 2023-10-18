import type { TSquare } from '$lib/types';

//api
export const DEFAULT_API_PATH = 'http://localhost:3000';
export const HEADERS_WITH_BEARER_TOKEN = (token: string, withJson: boolean = false) => {
	console.log(token)
	let tokenHeader = {
		'Authorization': `Bearer ${token}`,
	};

	let resultHeaders = withJson ? {...tokenHeader, 'Content-Type': 'application/json'} : tokenHeader;

	return {
		headers: resultHeaders
	};
};

export const initialSquareCoordinates: Omit<TSquare, 'id'> = {
	x_top: [0, 0],
	x_bottom: [0, 0],
	y_top: [0, 0],
	y_bottom: [0, 0]
};
