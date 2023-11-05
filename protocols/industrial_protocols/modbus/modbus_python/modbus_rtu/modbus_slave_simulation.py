from pymodbus.server import StartTcpServer
from pymodbus.datastore import ModbusSequentialDataBlock
from pymodbus.datastore import ModbusSlaveContext, ModbusServerContext

def run_modbus_server(address, port, slave_id):
    # Create a Modbus RTU slave context and initialize data blocks
    block = ModbusSequentialDataBlock(0, [0] * 100)  # Holding registers
    store = ModbusSlaveContext(hr=block)

    # Create a Modbus server context with the defined slave context
    context = ModbusServerContext(slaves=store, single=False)

    # Create and start a Modbus RTU TCP server
    server = StartTcpServer(context, address=(address, port))

if __name__ == "__main__":
    run_modbus_server("localhost", 5020, 1)  # Server 1
    run_modbus_server("localhost", 5021, 2)  # Server 2
