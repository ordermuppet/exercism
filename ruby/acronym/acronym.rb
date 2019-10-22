module Acronym
  def self.abbreviate(name)
    name.scan(/\b[a-zA-z]/).join.upcase
  end
end