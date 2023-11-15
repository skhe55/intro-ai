import { DEFAULT_API_PATH } from "$constants/index";
import type { TImage, TImageDeleteDto, TImageDto, TStandartApiResponse } from "$api/types";
import { customFetch } from "../fetchClient";

export class ImageApi {
    createImage = async (payload: TImageDto): Promise<TStandartApiResponse<string> | null> => {
        try {
            const response = await customFetch(`${DEFAULT_API_PATH}/images/create`, {
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
            console.error(`createImage: ${e}`);
            return null;
        }
    };
    
    uploadImage = async (id: string, projectId: string, file: File): Promise<TStandartApiResponse<string> | null> => {
        try {
            const formData = new FormData();
            formData.append("file", file);

            const response = await customFetch(`${DEFAULT_API_PATH}/images/upload/${id}?projectId=${projectId}`, {
                method: "POST",
                body: formData,
            });

            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`uploadImage: ${e}`);
            return null;
        }
    };

    getImages = async (projectId: string): Promise<TStandartApiResponse<TImage[]> | null> => {
        try {
            const response = await customFetch(`${DEFAULT_API_PATH}/images/all/${projectId}`);
            if (response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`getImages: ${e}`);
            return null;
        }
    };

    getImageById = async (imageId: string): Promise<TStandartApiResponse<TImage> | null> => {
        try {
            const response = await customFetch(`${DEFAULT_API_PATH}/images/${imageId}`);
            if (response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`getImageById: ${e}`);
            return null;
        }
    };

    deleteImage = async (imageId: string, dto: TImageDeleteDto): Promise<TStandartApiResponse<string> | null> => {
        try {
            const response = await customFetch(`${DEFAULT_API_PATH}/images/delete/${imageId}`, {
                method: "DELETE",
                headers: {
                    'Content-Type': "application/json"
                },
                body: JSON.stringify(dto)
            });
            if (response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`deleteImage: ${e}`);
            return null;
        }
    };
}