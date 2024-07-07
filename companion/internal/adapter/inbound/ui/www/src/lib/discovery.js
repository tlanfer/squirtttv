import {derived, readable} from "svelte/store";
import {config} from "./config.js";

export const setScanning = (value) => {
    scanning = value;
}

let scanning = false;

export const discovery = readable([], set =>{

    let update = async () => {
        if (scanning) {
            const response = await fetch("/api/discovery");
            let data = await response.json();
            set(data.hosts || []);
        }

        setTimeout(update, 1000)
    };
    setTimeout(update, 1);

    return () => {};
})

export const newHosts = derived([discovery, config], ([$discovery, $config], set) => {
    let newHosts = [];
    $discovery.forEach(host => {
        if (!$config.devices.find(device => device.host === host)) {
            newHosts.push(host);
        }
    });
    set(newHosts);
})