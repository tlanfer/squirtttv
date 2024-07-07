import {readable} from "svelte/store";

const initialState = {
    twitch: false,
    streamlabs: false,
    streamelements: false,
}

export const status = readable(initialState, set => {
    let state = initialState
    const getState = async () => {
        const response = await fetch("/api/status?wait=true" +
            "&twitch="+ state.twitch +
            "&streamlabs="+ state.streamlabs +
            "&streamelements=" + state.streamelements
        )
        state = await response.json();
        set(state);
        setTimeout(getState, 1000)
    }
    getState();
})