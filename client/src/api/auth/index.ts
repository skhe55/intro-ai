import { DEFAULT_API_PATH, JSON_BODY_HEADERS } from "$constants/index";
import type { TSignInPayload, TSignInUpResponse, TSignUpPayload, TStandartApiResponse } from "$api/types";

export class AuthApi {
    signIn = async (payload: TSignInPayload): Promise<TStandartApiResponse<TSignInUpResponse> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/auth/login`, {
                method: "POST",
                ...JSON_BODY_HEADERS,
                body: JSON.stringify(payload),
            });
            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`sign-in: ${e}`);
            return null;
        }
    };

    signUp = async (payload: TSignUpPayload): Promise<TStandartApiResponse<TSignInUpResponse> | null> => {
        try {
            const response = await fetch(`${DEFAULT_API_PATH}/auth/register`, {
                method: "POST",
                ...JSON_BODY_HEADERS,
                body: JSON.stringify(payload),
            });
            if(response.ok) {
                return response.json();
            } else {
                return null;
            }
        } catch(e) {
            console.error(`sign-up: ${e}`);
            return null;
        }
    };
}