function duplicateCount(text){
    const frequency = {}
    const lowerCaseText = text.toLowerCase()

    for (const character of lowerCaseText) {
        if (frequency[character])
            frequency[character] += 1
        else
            frequency[character] = 1
    }

    let answer = 0;
    for (const key in frequency) {
        if (frequency[key] > 1)
            ++answer
    }

    return answer
}