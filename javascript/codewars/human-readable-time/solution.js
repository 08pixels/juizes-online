function humanReadable (seconds) {

    const pad = (number) => {
        return number < 10 ? '0'+number : number
    }

    const minuteInSeconds = 60
    const hourInSeconds   = 60 * minuteInSeconds

    const HH = parseInt(seconds / hourInSeconds)
    const MM = parseInt((seconds % hourInSeconds) / minuteInSeconds)
    const SS = parseInt((seconds % hourInSeconds) % minuteInSeconds)

    const [hh,mm,ss] = [HH,MM,SS].map(pad)

    return `${hh}:${mm}:${ss}`
}