class Matrix
  attr_reader :rows

  def initialize(matrix)
    @rows = matrix.split("\n").map { |row| row.split.map(&:to_i) }
  end

  def columns
    rows.transpose
  end
end