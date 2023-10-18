export type TStandartApiResponse<T> = {
    Message: string;
    Status: string;
    Result: T;
}

export type TSignInPayload = {
    login: string;
    password: string;
};

export type TSignInUpResponse = {
    username: string;
    token: string;
    expires_at: string;
};

export type TSignUpPayload = TSignInPayload & { username?: string };

// Projects

export type TProject = {
    id: string;
    name: string;
    created_at: Date;
    updated_at: Date;
};

export type TProjectDto = {
    name: string;
};