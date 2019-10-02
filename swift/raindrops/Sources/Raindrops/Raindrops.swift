class Raindrops {
    var number: Int
    var factors: [Int] = []
    
    var sounds: String {
        get {
            var soundsString = ""
            if factors.contains(3) {
                soundsString.append("Pling")
            }
            if factors.contains(5) {
                soundsString.append("Plang")
            }
            if factors.contains(7) {
                soundsString.append("Plong")
            }
            if soundsString == "" {
                soundsString = String(number)
            }
            return soundsString
        }
    }

    init(_ number: Int) {
        self.number = number
        
        var low = 1, high = number
        factors += [low, high]
        while high - low > 1 {
            low += 1
            if number % low == 0 {
                high = number / low
                factors += [low, high]
            }
        }
    }
}
