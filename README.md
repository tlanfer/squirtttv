# Squirtianna

This lets you squirt someone with water when something happens on twitch.

## Building the hardware

1. Wire up an ESP8266 so that pin D6 triggers the water pistol

* Start with a "Gloria AutoPump Mini" electronic water sprayer
* Remove the original micro switch from the trigger
* Replace the micro switch with a mosfet, or some other way for GPIO 12 to trigger the pump.
* A neat solution is using a Wemos D1 mini and a mosfet shield for it

2. Install the firmware on it using Arduino

I'm sure this will work with other electric sprayers as well, but i only tested it with the Gloria AutoPump Mini.

## Getting the software

1. [Go here](/releases/tag/latest) and download the latest `companion.exe`
2. Run the executable once. It will create an example config file named `config.example.yaml`
3. Copy `config.example.yaml` to `config.yaml`
4. Customize the config to your liking
5. Run the executable

## Configuration reference

```yaml
cooldown: 5s
duration: 1s
twitch: arianna
streamlabs: eyJ0eX.... get yours from https://streamlabs.com/dashboard#/settings/api-settings > API Tokens > Your Socket API Token
events:
  - type: bits
    min: 0
    max: 100
  - type: bits
    min: 200
    max: 250
  - type: dono
    min: 20
    max: 30
  - type: dono
    min: 100
  - type: subs
    min: 10
```

| Key         | Description                                                                                                                                        |
|-------------|----------------------------------------------------------------------------------------------------------------------------------------------------|
| cooldown    | Ignore events after a spray happened                                                                                                               |
| duration    | How long to activate the sprayer. Around 500ms-1s seems to be alright.                                                                             |
| twitch      | Watch this channel for cheers or subs. You can leave this empty if you dont want twitch integration. You need at least one of twitch/streamlabs.   |
| streamlabs  | Connecto to streamlabs for donations. You can leave this empty if you dont want to react to donations. You need at least one of twitch/streamlabs. |
| events      | A list of events that can trigger the sprayer. If at least one matches, the sprayer activates.                                                     |
| events.type | One of "bits", "dono", "subs".                                                                                                                     |
| events.min  | Minimum amount for this event. Defaults to 0 if left out.                                                                                          |
| events.max  | Maximum amount for this event. Can be left out, defaults to infinity.                                                                              |

It should detect the sprayer on your network and connect to twich/streamlabs.