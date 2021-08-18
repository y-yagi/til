require 'ruby-watchman'
require 'socket'
require 'pathname'

sockname = RubyWatchman.load(
  %x{watchman --output-encoding=bser get-sockname}
)['sockname']
raise unless $?.exitstatus.zero?

UNIXSocket.open(sockname) do |socket|
  root = Pathname.new('/path/to/directory').realpath.to_s

  clock = RubyWatchman.query(['clock', root], socket)['clock']
  pp "current clock: #{clock}"

  paths = RubyWatchman.query(['since', root, ARGV[0]], socket)
  pp paths["files"].map { |f| f["name"] }
end
