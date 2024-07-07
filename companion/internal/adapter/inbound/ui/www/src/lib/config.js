import {derived, writable} from "svelte/store";

export const dirty = writable(false);
export const config = writable({
    devices: [],
    settings: {},
    events: {},
});
let debounce;

const load = async () => {
    const response = await fetch("/api/config");
    const data = await response.json();
    config.set(data);
}

const save = async (data) => {
    const response = await fetch("/api/config", {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(data)
    });
}

load();


config.subscribe((value) => {
    dirty.set(true);
    if (debounce) clearTimeout(debounce);
    debounce = setTimeout(async() => {
        await save(value)
        dirty.set(false);
    }, 1000);
});

export const squirters = derived(config, ($config) => {
    return $config.devices;
});

export const currency = derived(config, ($config) => {
    const baseCurrency = $config.settings.baseCurrency;

    if (baseCurrency === "usd") {
        return {
            unit: "p",
            units: "p"
        }
    }

    if (baseCurrency === "gbp") {
        return {
            unit: "p",
            units: "p"
        }
    }

    if (baseCurrency === "eur") {
        return {
            unit: "cent",
            units: "cents"
        }
    }

    return {
        unit: "p",
        units: "p"
    }
})