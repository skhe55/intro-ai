import type { TSquare } from "$lib/types";
import type { TAnnotationWithLabelNames } from "$api/types";

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

export const mappingAnnotationsToSquare = (dto: TAnnotationWithLabelNames[]): TSquare[] => {
    const result: TSquare[] = [];
    for(const item of dto) {
        result.push(
            {
                id: item.id,
                label_name: item.label_name,
                restored: true,
                x_top: item.coordinates[0] as [number, number],
                x_bottom: item.coordinates[1] as [number, number],
                y_top: item.coordinates[2] as [number, number],
                y_bottom: item.coordinates[3] as [number, number],
            }
        )
    }
    return result;
};

export const exportAnnotations = (data: TAnnotationWithLabelNames[]) => {
    if(!data.length) {
        return null;
    } else {
        const result = [];
        const columns = Object.keys(data[0]);
        for(let i = 0; i < data.length+1; i++) {
            if (i === 0) {
                result.push(columns);
            } else {
                result.push(
                    [
                        data[i-1].id,
                        data[i-1].label_id,
                        data[i-1].label_name,
                        data[i-1].coordinates,
                        data[i-1].created_at,
                    ]
                )
            }
        }
        const csvContent = "data:text/csv;charset=utf-8," + result.map(e => e.join(",")).join("\n");
        const encodedUri = encodeURI(csvContent);
        const link = document.createElement("a");
        link.setAttribute("href", encodedUri);
        link.setAttribute("download", "my_data.csv");
        document.body.appendChild(link); // Required for FF

        link.click(); 
    }
};