module Complement
  def self.of_dna(dna_string)
    nucleotides = dna_string.split('')
    return '' unless nucleotides.all? { |nucleotide| nucleotide_mapping.keys.include?(nucleotide) }
    nucleotides.map { |nucleotide| nucleotide_mapping[nucleotide] }.join
  end

  def self.nucleotide_mapping
    {
      'G' => 'C',
      'C' => 'G',
      'T' => 'A',
      'A' => 'U'
    }
  end
end

module BookKeeping
  VERSION = 4
end
