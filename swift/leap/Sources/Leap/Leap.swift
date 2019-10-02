class Year {
    var calendarYear: Int
    var isLeapYear: Bool {
        get {
            return divisibleBy(400) || (divisibleBy(4) && !divisibleBy(100))
        }
    }
    
    init(calendarYear: Int) {
        self.calendarYear = calendarYear
    }
    
    private func divisibleBy(_ number: Int) -> Bool {
        return calendarYear % number == 0
    }
}
