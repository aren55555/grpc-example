this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
puts lib_dir
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)


require 'grpc'
require 'multi_json'
require 'garage_services_pb'

def request_car(n)
  stub = Garage::Garage::Stub.new('localhost:50051', :this_channel_is_insecure)
  car = stub.random_car(Garage::CarRequest.new(requestNumber: n))
  puts "A random car was retrieved from the GRPC server:" +
         "\n\tMake: #{car.make}" +
         "\n\tModel: #{car.model}" +
         "\n\tYear: #{car.year}\n\t" +
         "HP: #{car.horsepower}"
end

def main
  request_number = 1
  while true do
    request_car(request_number)
    request_number += 1
    sleep 1
  end
end


main() # start client app
