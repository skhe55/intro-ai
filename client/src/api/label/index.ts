import { DEFAULT_API_PATH } from "$constants/index";
import type { TLabel, TLabelDto, TStandartApiResponse } from "$api/types";

export class LabelApi {
    createLabel = async (payload: TLabelDto): Promise<TStandartApiResponse<string> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/labels/create`, {
                method: "POST",
                headers: {
                    'Authorization': `Bearer ${window.localStorage.getItem('token') as string}`,
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
            console.error(`createLabel: ${e}`);
            return null;
        }
    };

    getLabelsByImageId = async (imageId: string): Promise<TStandartApiResponse<TLabel[]> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/labels?imageId=${imageId}`, {
                headers: {
                    'Authorization': `Bearer ${window.localStorage.getItem('token')}`
                },
            });
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
    
    deleteLabel = async (id: string): Promise<TStandartApiResponse<string> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/labels/delete/${id}`, {
                method: "DELETE",
                headers: {
                    'Authorization': `Bearer ${window.localStorage.getItem('token') as string}`,
                },
            });

            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`deleteLabel: ${e}`);
            return null;
        }
    };
}