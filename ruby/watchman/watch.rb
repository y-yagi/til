require 'ruby-watchman'
require 'socket'
require 'pathname'

sockname = RubyWatchman.load(
  %x{watchman --output-encoding=bser get-sockname}
)['sockname']
raise unless $?.exitstatus.zero?

UNIXSocket.open(sockname) do |socket|
  root = Pathname.new('/path/to/watch').realpath.to_s
  roots = RubyWatchman.query(['watch-list'], socket)['roots']
  if !roots.include?(root)
    result = RubyWatchman.query(['watch', root], socket)
    raise if result.has_key?('error')
  end
end
