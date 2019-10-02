class Integer
  def to_roman
    RomanNumeral.create_numerals(self).map(&:to_string).join('')
  end
end

class RomanNumeral
  attr_reader :digit, :one, :five, :ten

  def initialize(digit)
    @digit = digit
    initialize_numerals
  end

  def self.create_numerals(integer)
    places = [Thousands, Hundreds, Tens, Ones]

    # zero-padded array of digits, i.e. 911 becomes [0, 9, 1, 1]
    digits = ("%.#{places.size}d" % integer).chars.map(&:to_i)

    digits.each_with_index.map { |digit, i| places[i].new(digit) }
  end

  def to_string
    ['', one, one * 2, one * 3, one + five, five, five + one, five + (one * 2), five + (one * 3), one + ten][digit]
  end

  def initialize_numerals
    fail NotImplemented
  end
end

class Thousands < RomanNumeral
  def initialize_numerals
    @one = 'M'
  end

  def to_string
    one * digit
  end
end

class Hundreds < RomanNumeral
  def initialize_numerals
    @one = 'C'
    @five = 'D'
    @ten = 'M'
  end
end

class Tens < RomanNumeral
  def initialize_numerals
    @one = 'X'
    @five = 'L'
    @ten = 'C'
  end
end

class Ones < RomanNumeral
  def initialize_numerals
    @one = 'I'
    @five = 'V'
    @ten = 'X'
  end
end

module BookKeeping
  VERSION = 2
end
