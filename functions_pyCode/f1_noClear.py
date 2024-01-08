# importing the requests library
import requests
import json
 
# api-endpoint
log_url = "http://localhost:8080"

db_address = "localhost:6379"
 
# defining a params dict for the parameters to be sent to the API
# PARAMS = {'db_address': 'localhost:50000', 'key': 'f1_x', 'value': 0, 'ssf_id': 123456, 'step_id': 1}
 
# sending get request and saving the response as response object
r = requests.post(url = log_url + '/read', 
                  params={'db_address': db_address, 
                          'key': 'f1_x', 
                        #   'value': 0, 
                          'ssf_id': 123456, 
                          'step_id': 1})
print(json.dumps(r.json()))

r = requests.post(url = log_url + '/write', 
                  params={'db_address': db_address, 
                          'key': 'f1_x', 
                          'value': 23, 
                          'ssf_id': 123456, 
                          'step_id': 2})
print(json.dumps(r.json()))

r = requests.post(url = log_url + '/read', 
                  params={'db_address': db_address, 
                          'key': 'f1_x', 
                        #   'value': 0, 
                          'ssf_id': 123456, 
                          'step_id': 2})
print(json.dumps(r.json()))

r = requests.post(url = log_url + '/write', 
                  params={'db_address': db_address, 
                          'key': 'f1_x', 
                          'value': 233, 
                          'ssf_id': 123456, 
                          'step_id': 3})
print(json.dumps(r.json()))


