import type { TSquare } from "$lib/types";

export const getPointsUpperRightCorner = (coordinates: TSquare) => {
    if(coordinates.y_top[1] < coordinates.y_bottom[1]) {
        return [
            Math.max(coordinates.y_top[0], coordinates.x_top[0]) + 5, 
            Math.max(coordinates.y_top[1], coordinates.x_top[1]), 
            Math.max(coordinates.y_top[0], coordinates.x_top[0]) + 25, 
            Math.max(coordinates.y_top[1], coordinates.x_top[1]) + 20
        ];
    } else {
        return [
            Math.max(coordinates.y_bottom[0], coordinates.x_bottom[0]) + 5, 
            Math.max(coordinates.y_bottom[1], coordinates.x_bottom[1]), 
            Math.max(coordinates.y_bottom[0], coordinates.x_bottom[0]) + 25, 
            Math.max(coordinates.y_bottom[1], coordinates.x_bottom[1]) + 20
        ];
    }
};