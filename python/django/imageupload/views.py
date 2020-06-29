from django.apps import apps
from django.core.files.uploadedfile import UploadedFile
from django.http import HttpResponse, HttpRequest, JsonResponse


# Create your views here.
from django.views.decorators.csrf import csrf_exempt
from django.views.decorators.http import require_POST


@csrf_exempt
@require_POST
def upload(request: HttpRequest) -> HttpResponse:
    config = apps.get_app_config('imageupload')
    if config.form_image_key not in request.FILES:
        return JsonResponse({'message': 'missing file'}, status=400)
    handle_file(request.FILES[config.form_image_key])
    return HttpResponse(status=204)


def handle_file(file: UploadedFile):
    # do something with the file
    pass
