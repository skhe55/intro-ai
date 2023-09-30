// canvas
export type TPoint = [number, number];
export type TDrawFn = (ctx: CanvasRenderingContext2D) => void;
export type TCanvasContext = {
	addDrawFn: (item: TCanvasItem) => void;
	removeDrawFn: (item: TCanvasItem) => void;
};
export type TCanvasItem = {
	id: string;
	fn: TDrawFn;
};

// markup page
export type TSquare = {
	id: string;
	x_top: TPoint;
	x_bottom: TPoint;
	y_top: TPoint;
	y_bottom: TPoint;
};
