class VehicleModel {
  int id;
  int year;
  String model;
  String plate;

  VehicleModel({
    required this.id,
    required this.year,
    required this.model,
    required this.plate,
  });

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'id': id,
      'year': year,
      'model': model,
      'plate': plate,
    };
  }

  factory VehicleModel.fromJson(Map<String, dynamic> json) {
    return VehicleModel(
      id: json['id'] as int,
      year: json['year'] as int,
      model: json['model'] as String,
      plate: json['plate'] as String,
    );
  }
}
