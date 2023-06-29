import pika
import uuid
import time
import task_pb2

class RPCClient:
    def __init__(self):
        self.connection = pika.BlockingConnection(pika.ConnectionParameters('localhost'))
        self.channel = self.connection.channel()

        result = self.channel.queue_declare(queue='', exclusive=True)
        self.callback_queue = result.method.queue

        self.channel.basic_consume(queue=self.callback_queue, on_message_callback=self.on_response, auto_ack=True)

    def on_response(self, ch, method, props, body):
        if self.correlation_id == props.correlation_id:
            self.response = body

    def call(self, tasks):
        self.response = None
        self.correlation_id = str(uuid.uuid4())

        task_input_array = task_pb2.TaskInputArray()
        for task in tasks:
            task_input = task_input_array.tasks.add()
            task_input.wbs = task["wbs"]
            task_input.wbs_parent = task["wbs_parent"]
            task_input.wbs_children.extend(task["wbs_children"])
            task_input.status_date = task["status_date"]
            task_input.duration_statusdate = task["duration_statusdate"]
            task_input.cost = task["cost"]
            task_input.progress = task["progress"]
            task_input.baseline_start_date = task["baseline_start_date"]
            task_input.baseline_end_date = task["baseline_end_date"]
            task_input.baseline_duration = task["baseline_duration"]
            task_input.baseline_cost = task["baseline_cost"]
            task_input.BCWP = task["BCWP"]
            task_input.BCWS = task["BCWS"]
            task_input.SPI = task["SPI"]
            task_input.percent_plan = task["percent_plan"]
            task_input.percent_actual = task["percent_actual"]

        body = task_input_array.SerializeToString()

        self.channel.basic_publish(
            exchange='',
            routing_key='rpc_queue',
            properties=pika.BasicProperties(
                reply_to=self.callback_queue,
                correlation_id=self.correlation_id,
                content_type='application/protobuf',
            ),
            body=body)

        start_time = time.time()

        while self.response is None:
            self.connection.process_data_events()

        processing_time = time.time() - start_time

        print(f"Processing time: {processing_time} seconds")

        return self.response

def generate_mockup_tasks():
    # Generate mockup tasks
    tasks = [
        {
            "wbs": "1",
            "wbs_parent": "",
            "wbs_children": ["1.1", "1.2"],
            "status_date": "2023-06-22T00:00:00",
            "duration_statusdate": 30,
            "cost": 10000.0,
            "progress": 0.5,
            "baseline_start_date": "2023-06-01T00:00:00",
            "baseline_end_date": "2023-08-01T00:00:00",
            "baseline_duration": 60,
            "baseline_cost": 20000.0,
            "BCWP": 10000.0,
            "BCWS": 15000.0,
            "SPI": 0.67,
            "percent_plan": 0.5,
            "percent_actual": 0.67
        },
        {
            "wbs": "1.1",
            "wbs_parent": "1",
            "wbs_children": [],
            "status_date": "2023-06-22T00:00:00",
            "duration_statusdate": 15,
            "cost": 5000.0,
            "progress": 0.75,
            "baseline_start_date": "2023-06-01T00:00:00",
            "baseline_end_date": "2023-07-01T00:00:00",
            "baseline_duration": 30,
            "baseline_cost": 10000.0,
            "BCWP": 7500.0,
            "BCWS": 10000.0,
            "SPI": 0.75,
            "percent_plan": 0.5,
            "percent_actual": 0.75
        },
        {
            "wbs": "1.2",
            "wbs_parent": "1",
            "wbs_children": [],
            "status_date": "2023-06-22T00:00:00",
            "duration_statusdate": 20,
            "cost": 6000.0,
            "progress": 0.6,
            "baseline_start_date": "2023-06-02T00:00:00",
            "baseline_end_date": "2023-07-02T00:00:00",
            "baseline_duration": 30,
            "baseline_cost": 10000.0,
            "BCWP": 6000.0,
            "BCWS": 9000.0,
            "SPI": 0.66,
            "percent_plan": 0.5,
            "percent_actual": 0.6
        },
    ]
    return tasks

def main():
    rpc_client = RPCClient()

    tasks = generate_mockup_tasks()
    response = rpc_client.call(tasks)
    task_output_array = task_pb2.TaskOutputArray()
    task_output_array.ParseFromString(response)

    for task_output in task_output_array.tasks:
        print("Task:")
        print(f"WBS: {task_output.wbs}")
        print(f"Progress: {task_output.progress}")
        print(f"Cost: {task_output.cost}")
        print(f"BCWP: {task_output.BCWP}")
        print(f"BCWS: {task_output.BCWS}")
        print(f"SPI: {task_output.SPI}")
        print(f"Percent Plan: {task_output.percent_plan}")
        print(f"Percent Actual: {task_output.percent_actual}")
        print()

if __name__ == '__main__':
    main()
