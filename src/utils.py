import logging

from redis import Redis
from rq import Queue

from .config import REDIS_HOST, REDIS_PORT, REDIS_DB

logger = logging.getLogger(__name__)

def get_redis_client():
    """Returns a Redis client instance."""
    return Redis(host=REDIS_HOST, port=REDIS_PORT, db=REDIS_DB, decode_responses=True)

def get_task_queue():
    """Returns a Redis queue instance."""
    return Queue(connection=get_redis_client())

def get_redis_connection():
    """Returns a Redis connection instance."""
    return get_redis_client().connection