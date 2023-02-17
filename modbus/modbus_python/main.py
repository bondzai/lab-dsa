from pymodbus.client import ModbusTcpClient
import struct

# Connect to Modbus TCP server
client = ModbusTcpClient('192.168.11.87', port=502)
client.connect()

# Read holding registers starting from address 0
response = client.read_holding_registers(address=99, count=4, slave=1)

# Print the register values for debugging
print(response.registers)

# Combine the four registers into a single floating point value
if len(response.registers) == 4:
    byte_string = struct.pack('>HHHH', response.registers[0], response.registers[1], response.registers[2], response.registers[3])
    print(byte_string)
    if len(byte_string) == 4:
        float_value = struct.unpack('>f', byte_string)[0]
        print(float_value)
    else:
        print("Error: byte string length is not 4")
else:
    print("Error: response does not contain 4 registers")

# Close the connection
client.close()