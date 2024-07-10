import {readable} from "svelte/store";

export const version = readable({version: "", latest:"", isLatest: true}, async set => {
    const res = await fetch("/api/version")
    const ver = await res.json()

    set(ver)
});

export const runUpdate = async ()=> {
    await fetch("/api/version", {
        method: "POST"
    })
}