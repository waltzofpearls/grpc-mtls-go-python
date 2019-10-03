import api.metrics_pb2
import api.metrics_pb2_grpc
import grpc
import os
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.duration_pb2 import Duration
from datetime import datetime, timedelta

class Client:
    def __init__(self):
        self.root_ca = os.environ.get('TLS_ROOT_CA', '').encode()
        self.tls_key = os.environ.get('TLS_CLIENT_KEY', '').encode()
        self.tls_cert = os.environ.get('TLS_CLIENT_CERT', '').encode()
        self.address = os.environ.get('GRPC_SERVER_ADDRESS', '')

    def fetch(self):
        try:
            credentials = grpc.ssl_channel_credentials(
                self.root_ca, self.tls_key, self.tls_cert)
            channel = grpc.secure_channel(self.address, credentials)
            stub = api.metrics_pb2_grpc.MetricsStub(channel)

            end_dt = datetime.now()
            start_dt = end_dt - timedelta(hours=24)

            start_ts, end_ts = Timestamp(), Timestamp()
            start_ts.FromDatetime(start_dt)
            end_ts.FromDatetime(end_dt)

            req = api.metrics_pb2.QueryMetricsRequest(
                metricName='steps',
                start=start_ts,
                end=end_ts
            )
            resp = stub.Query(req)
            return resp
        except Exception as e:
            print('ERROR:', e)
            return None

if __name__ == "__main__":
    c = Client()
    metrics = c.fetch()
    print(metrics)
