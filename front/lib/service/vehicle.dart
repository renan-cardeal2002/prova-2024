import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:front/model/vehicle.dart';

class VehicleService {
  Future<http.Response> postData({required int year, required String model, required String plate}) async {
    try {
      final response = await http.post(
        Uri.parse('http://localhost:8888/api/vehicle'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, dynamic>{'year': year, 'model': model, 'plate': plate}),
      );

      return response;
    } catch (e) {
      throw Exception('Error: $e');
    }
  }

  Future<List<VehicleModel>> fetchData() async {
    try {
      final request = await http.get(
        Uri.parse('http://localhost:8888/api/vehicle'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
      );

      if (request.statusCode == 200) {
        final Map<String, dynamic> jsonBody = jsonDecode(request.body);
        var list = jsonBody['data'] as List;
        List<VehicleModel> vehicles = list.map((u) => VehicleModel.fromJson(u)).toList();

        return vehicles;
      } else {
        throw Exception('Failed to load vehicle data');
      }
    } catch (e) {
      throw Exception('Erro na requisição $e');
    }
  }

  Future<void> updateData({required int id, required int year, required String model, required String plate}) async {
    try {
      final response = await http.patch(
        Uri.parse('http://localhost:8888/api/vehicle/$id'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
        body: jsonEncode(<String, String>{'year': year, 'model': model, 'plate': plate}),
      );
    } catch (e) {
      throw Exception('Erro na requisição $e');
    }
  }

  Future<void> deleteData({required int id}) async {
    try {
      final response = await http.delete(
        Uri.parse('http://localhost:8888/api/vehicle/$id'),
        headers: <String, String>{
          'Content-Type': 'application/json; charset=UTF-8',
        },
      );
    } catch (e) {
      throw Exception('Erro na requisição $e');
    }
  }
}
