import pytest
import requests
import os

def test_basic():
    assert 1 == 1

def test_url_basic():
    # read env for CLAIMSAI_RESOLVER_LOCAL_NAME and set it to localhost if empty
    HOST_NAME = os.getenv('CLAIMSAI_RESOLVER_LOCAL_NAME', "localhost")     
    url = f"http://{HOST_NAME}:8889/"
    r = requests.get(url)
    assert r.status_code == 200