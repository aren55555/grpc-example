this_dir = File.expand_path(File.dirname(__FILE__))
lib_dir = File.join(this_dir, 'lib')
puts lib_dir
$LOAD_PATH.unshift(lib_dir) unless $LOAD_PATH.include?(lib_dir)


require 'grpc'
require 'multi_json'
require 'parking_services_pb'

GRPC_SERVER = 'localhost:50051'
# GRPC_SERVER = 'grpc.arenpatel.com:50051'

def request_status(stub)
  status = stub.get_status(Parking::GetStatusRequest.new())
  puts status
end

def park_car(stub)
  response = stub.park_car(Parking::ParkCarRequest.new(
    car: Parking::Car.new(
      make: 'Subaru',
      model: 'WRX',
      year: 2019,
      horsepower: 268
    )
  ))
  puts response
end

def main
  stub = Parking::ParkingService::Stub.new(GRPC_SERVER, :this_channel_is_insecure)

  park_car(stub)

  # while true do
  #   request_status(stub)
  #   sleep(5)
  # end
end


main() # start client app
