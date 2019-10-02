# A class for computing the number of differences in two strings
# (Hamming distance)
class Hamming
  def self.compute(first_string, second_string)
    raise ArgumentError if first_string.length != second_string.length

    (0...first_string.length).count { |i| first_string[i] != second_string[i] }
  end
end

module BookKeeping
  VERSION = 3
end