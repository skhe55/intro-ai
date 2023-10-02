export const navigate = (path: string): void => {
    const a = document.createElement("a");
    a.setAttribute("href", `/${path}`);
    a.click();
    document.removeChild(a);
};

export const isTokenExpires = (date: string | null): boolean => {
    if(!date) {
        return true;
    } 
    const expiresAtDate = new Date(date).getTime();
    const dateNow = new Date().getTime();
    if(expiresAtDate > dateNow) {
        return false;
    } else {
        return true;
    }
};