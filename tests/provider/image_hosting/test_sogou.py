import os.path
import tempfile

from PIL import Image

from toolkit.provider.image_hosting import sogou


def test_sogou_upload():
    with tempfile.TemporaryDirectory() as tmpdir:
        file = os.path.join(tmpdir, "test.png")
        Image.new("RGB", (100, 100)).save(file)
        result = sogou.upload(file)
        assert result.success
        assert result.remote_url
