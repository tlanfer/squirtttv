import {readable} from "svelte/store";

export const version = readable({version: ""}, async set => {
    const res = await fetch("/api/version")
    const ver = await res.json()
    set(ver)
});
