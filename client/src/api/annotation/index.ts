import { DEFAULT_API_PATH } from "$constants/index";
import type { TAnnotation, TAnnotationDto, TAnnotationWithLabelNames, TStandartApiResponse } from "$api/types";
import { customFetch } from "../fetchClient";

export class AnnotationApi {
    createAnnotation = async (payload: TAnnotationDto): Promise<TStandartApiResponse<string> | null> => {
        try {
            const response = await customFetch(`${DEFAULT_API_PATH}/annotations/create`, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(payload),
            });
            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`createAnnotation: ${e}`);
            return null;
        }
    };

    getAnnotationByImageId = async (imageId: string): Promise<TStandartApiResponse<TAnnotationWithLabelNames[]> | null> => {
        try {
            const response = await customFetch(`${DEFAULT_API_PATH}/annotations?imageId=${imageId}`);
            if (response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`getAnnotationByImageId: ${e}`);
            return null;
        }
    };
    
    deleteAnnotation = async (id: string): Promise<TStandartApiResponse<string> | null> => {
        try {
            const response = await customFetch(`${DEFAULT_API_PATH}/annotations/delete/${id}`, {
                method: "DELETE",
            });

            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`deleteAnnotation: ${e}`);
            return null;
        }
    };
}