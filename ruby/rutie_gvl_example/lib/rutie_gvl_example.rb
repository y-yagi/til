require 'rutie_gvl_example/version'
require 'rutie'

module RutieGvlExample
  Rutie.new(:rutie_gvl_example).init 'Init_rutie_gvl_example', __dir__
end

class RutieExample
  def self.watch(path, recursive: false, &block)
    Thread.new do
      RutieExample.raw_watch(path, recursive: recursive, &block)
    end
  end
end
