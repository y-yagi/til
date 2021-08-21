require 'ruby-watchman'
require 'socket'
require 'pathname'

sockname = RubyWatchman.load(
  %x{watchman --output-encoding=bser get-sockname}
)['sockname']
raise unless $?.exitstatus.zero?

UNIXSocket.open(sockname) do |socket|
  root = Pathname.new("/path/to/directory").realpath.to_s
  result = RubyWatchman.query(['subscribe', root, "mysubscription", { "fields" => ["name"], "expression" => ["allof", ["match", "*.rb"]] } ], socket)
  raise result.inspect if result.include?("error")

  loop do
    msg = socket.recvmsg()[0]
    pp RubyWatchman.load(msg)
  end
end
