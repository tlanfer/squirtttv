
export const test = async (pattern, choose, devices) => {

    await fetch('/api/test', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({
            pattern: pattern,
            choose: choose,
            devices: devices
        }),
    })

}