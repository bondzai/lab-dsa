def generate_mockup_tasks():
    tasks = [
        {
            "wbs": "1.1",
            "wbs_parent": "1",
            "wbs_children": ["1.1.1", "1.1.2"],
            "status_date": "2023-06-27",
            "actual_start_date": "2023-06-01",
            "actual_end_date": "2023-06-26",
            "duration_statusdate": 25,
            "cost": 10000,
            "progress": 95,
            "baseline_start_date": "2023-06-01",
            "baseline_end_date": "2023-06-25",
            "baseline_duration": 24,
            "baseline_cost": 9500,
            "baseline_version": 1.0,
            "BCWP": 9500,
            "BCWS": 9600,
            "SPI": 0.99,
            "percent_plan": 96,
            "percent_actual": 95,
        },
        # add more tasks here...
    ]
    return tasks
